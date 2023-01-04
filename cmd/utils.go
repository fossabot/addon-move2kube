package main

import (
	"fmt"

	"github.com/konveyor/tackle2-addon/command"
	"github.com/konveyor/tackle2-addon/repository"
	hub "github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
)

// addTags ensure tags created and associated with application.
// Ensure tag exists and associated with the application.
func addTags(application *api.Application, names ...string) error {
	addon.Activity("Adding tags: %v", names)
	appTags := appTags(application)
	// Fetch tags and tag types.
	tpMap, err := tpMap()
	if err != nil {
		return err
	}
	tagMap, err := tagMap()
	if err != nil {
		return err
	}
	// Ensure type exists.
	wanted := api.TagType{
		Name:  "DIRECTORY",
		Color: "#2b9af3",
		Rank:  3,
	}
	tp, found := tpMap[wanted.Name]
	if !found {
		tp = wanted
		if err := addon.TagType.Create(&tp); err != nil {
			return err
		}
		tpMap[tp.Name] = tp
	} else {
		if wanted.Rank != tp.Rank || wanted.Color != tp.Color {
			return &hub.SoftError{Reason: "Tag (TYPE) conflict detected."}
		}
	}
	// Add tags.
	for _, name := range names {
		if _, found := appTags[name]; found {
			continue
		}
		wanted := api.Tag{
			Name:    name,
			TagType: api.Ref{ID: tp.ID},
		}
		tg, found := tagMap[wanted.Name]
		if !found {
			tg = wanted
			if err := addon.Tag.Create(&tg); err != nil {
				return err
			}
			tagMap[wanted.Name] = tg
		} else {
			if wanted.TagType.ID != tg.TagType.ID {
				return &hub.SoftError{Reason: "Tag conflict detected."}
			}
		}
		addon.Activity("[TAG] Associated: %s.", tg.Name)
		application.Tags = append(
			application.Tags,
			api.Ref{ID: tg.ID},
		)
	}
	// Update application.
	return addon.Application.Update(application)
}

// tagMap builds a map of tags by name.
func tagMap() (map[string]api.Tag, error) {
	list, err := addon.Tag.List()
	if err != nil {
		return nil, err
	}
	m := map[string]api.Tag{}
	for _, tag := range list {
		m[tag.Name] = tag
	}
	return m, nil
}

// tpMap builds a map of tag types by name.
func tpMap() (map[string]api.TagType, error) {
	list, err := addon.TagType.List()
	if err != nil {
		return nil, err
	}
	m := map[string]api.TagType{}
	for _, t := range list {
		m[t.Name] = t
	}
	return m, nil
}

// appTags builds map of associated tags.
func appTags(application *api.Application) map[string]uint {
	m := map[string]uint{}
	for _, ref := range application.Tags {
		m[ref.Name] = ref.ID
	}
	return m
}

// commitResources commits the resources to the Git repo.
// func commitResources(SourceDir, groupId, artifactId string) error {
func commitResources(repo repository.Repository, inputDir, outputDir string) error {
	// Copy the output to into the repo
	cmd := command.Command{
		Path:    "/usr/bin/cp",
		Options: []string{"-r", outputDir, "move2kube-output"},
		Dir:     inputDir,
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy the output to the repo directory. Error: %w", err)
	}

	if err := repo.Branch("move2kube-output"); err != nil {
		return fmt.Errorf("failed to switch to a new branch. Error: %w", err)
	}
	if err := repo.Commit([]string{"-A"}, "feat: add move2kube transform output"); err != nil {
		return fmt.Errorf("failed to commit and push all the files. Error: %w", err)
	}

	// Copy the k8s resources to the output directory
	// cmd := command.Command{
	// 	Path: "/usr/bin/cp",
	// 	Options: []string{
	// 		"-r",
	// 		pathlib.Join(SourceDir, "target", "classes", "META-INF", "jkube/"), ".",
	// 		"output-------------------------------------------------",
	// 	},
	// 	Dir: SourceDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to copying resources. Error: %w", err)
	// }

	// Copy the Dockerfile to the output directory
	// group := strings.ToLower(strings.Split(groupId, ".")[1])
	// artifactId = strings.ToLower(artifactId)
	// cmd = command.Command{
	// 	Path: "/usr/bin/cp",
	// 	Options: []string{
	// 		pathlib.Join(SourceDir, "target", "docker", group, artifactId, "latest", "build", "Dockerfile"),
	// 		pathlib.Join("output-------------------------------------------", "Dockerfile")},
	// 	Dir: SourceDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to copying Dockerfile. Error: %w", err)
	// }

	// repG, ok := rep.(respository.Git)
	// g := repository.Git{}
	// addon.Activity("", g)

	// cmd := command.Command{
	// 	Path:    "/usr/bin/git",
	// 	Options: []string{"config", "--global", "user.email", "tackle@konveyor.org"},
	// 	Dir:     outputDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to set git config. Error: %w", err)
	// }

	// cmd = command.Command{
	// 	Path:    "/usr/bin/git",
	// 	Options: []string{"config", "--global", "user.name", "tackle"},
	// 	Dir:     outputDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to set git config. Error: %w", err)
	// }

	// cmd = command.Command{
	// 	Path:    "/usr/bin/git",
	// 	Options: []string{"add", pathlib.Base("output-------------------------------------------")},
	// 	Dir:     outputDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to adding k8s resources to git. Error: %w", err)
	// }

	// cmd = command.Command{
	// 	Path:    "/usr/bin/git",
	// 	Options: []string{"commit", "-m", "Add k8s resources"},
	// 	Dir:     outputDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to committing k8s resources. Error: %w", err)
	// }

	// cmd = command.Command{
	// 	Path:    "/usr/bin/git",
	// 	Options: []string{"push"},
	// 	Dir:     outputDir,
	// }
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to pushing k8s resources. Error: %w", err)
	// }

	return nil
}
