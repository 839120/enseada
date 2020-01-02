// Copyright 2019 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package maven

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/chartmuseum/storage"

	"github.com/enseadaio/enseada/internal/couch"
)

const StoragePrefix = couch.MavenDB

type RepoFile struct {
	Repo     *Repo
	Filename string
	Version  string
	content  []byte
	path     string
	storage  storage.Backend
}

func (f *RepoFile) Content() ([]byte, error) {
	if f.content == nil {
		obj, err := f.storage.GetObject(f.path)
		if err != nil {
			return nil, err
		}
		f.content = obj.Content
	}
	return f.content, nil
}
func (m *Maven) GetFile(ctx context.Context, path string) (*RepoFile, error) {
	m.Logger.Infof("looking up file with path %s", fmt.Sprintf(`"%s"`, path))
	repo, err := m.FindRepo(ctx, couch.Query{
		"files": couch.Query{
			"$elemMatch": couch.Query{
				"$eq": path,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if repo == nil {
		return nil, nil
	}

	slices := strings.Split(path, "/")
	return &RepoFile{
		Repo:     repo,
		Filename: slices[len(slices)-1],
		path:     path,
		storage:  m.storage,
	}, nil
}

func (m *Maven) PutFile(ctx context.Context, path string, content []byte) error {
	return m.storage.PutObject(path, content)
}

func (m *Maven) PutFileInRepo(ctx context.Context, repo *Repo, path string, content []byte) (*RepoFile, error) {
	trimmed := strings.TrimPrefix(path, repo.StoragePath)
	trimmed = strings.TrimPrefix(trimmed, "/")
	slices := strings.Split(trimmed, "/")
	filename := slices[len(slices)-1]
	var version string
	if len(slices) == 2 {
		version = slices[0]
	}
	file := &RepoFile{
		Repo:     repo,
		Filename: filename,
		Version:  version,
		content:  content,
	}
	m.Logger.Infof("storing file %+v", file)
	spath := filePath(file)
	err := m.PutFile(ctx, spath, content)
	if err != nil {
		return nil, err
	}

	repo.Files = append(repo.Files, spath)
	in := repo.Files
	sort.Strings(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[j] = in[i]
	}
	repo.Files = in[:j+1]
	return file, m.SaveRepo(ctx, repo)
}

func (m *Maven) ClearRepoStorage(ctx context.Context, repo *Repo) error {
	prefix := fmt.Sprintf("%s/%s/", StoragePrefix, repo.StoragePath)
	objs, err := m.storage.ListObjects(prefix)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.storage.DeleteObject(prefix + obj.Path); err != nil {
			return err
		}
	}

	repo.Files = []string{}
	return m.SaveRepo(ctx, repo)
}

func filePath(file *RepoFile) string {
	repo := file.Repo
	if file.Version == "" {
		return fmt.Sprintf("%s/%s/%s", StoragePrefix, repo.StoragePath, file.Filename)
	} else {
		return fmt.Sprintf("%s/%s/%s/%s", StoragePrefix, repo.StoragePath, file.Version, file.Filename)
	}
}
