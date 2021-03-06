package common

import (
	"fmt"
	"sort"

	"github.com/flant/werf/pkg/config"
	"github.com/flant/werf/pkg/storage"

	"github.com/flant/werf/pkg/util"
)

func GetManagedImageName(userSpecifiedImageName string) string {
	switch userSpecifiedImageName {
	case "~", storage.NamelessImageRecordTag:
		return ""
	}
	return userSpecifiedImageName
}

func GetManagedImagesNames(projectName string, stagesStorage storage.StagesStorage, werfConfig *config.WerfConfig) ([]string, error) {
	var imagesNames []string
	if managedImages, err := stagesStorage.GetManagedImages(projectName); err != nil {
		return nil, fmt.Errorf("unable to get managed images for project %q: %s", projectName, err)
	} else {
		imagesNames = append(imagesNames, managedImages...)
	}
	for _, image := range werfConfig.StapelImages {
		imagesNames = append(imagesNames, image.Name)
	}
	for _, image := range werfConfig.ImagesFromDockerfile {
		imagesNames = append(imagesNames, image.Name)
	}
	uniqImagesNames := util.UniqStrings(imagesNames)
	sort.Strings(uniqImagesNames)

	return uniqImagesNames, nil
}
