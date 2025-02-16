package config

import (
	"github.com/setlog/trivrost/pkg/system"
)

func FilterLauncherUpdatesByPlatform(launcherUpdates []LauncherUpdateConfig, os string, arch string) []LauncherUpdateConfig {
	var filteredLauncherUpdates []LauncherUpdateConfig
	for _, launcherUpdate := range launcherUpdates {
		if matchPlatform(launcherUpdate.TargetPlatforms, os, arch) {
			filteredLauncherUpdates = append(filteredLauncherUpdates, launcherUpdate)
		}
	}
	return filteredLauncherUpdates
}

func FilterBundlesByPlatform(bundles []BundleConfig, os string, arch string) []BundleConfig {
	var filteredBundles []BundleConfig
	for _, bundle := range bundles {
		if matchPlatform(bundle.TargetPlatforms, os, arch) {
			filteredBundles = append(filteredBundles, bundle)
		}
	}
	return filteredBundles
}

func FilterCommandsByPlatform(commands []Command, os string, arch string) []Command {
	var filteredCommands []Command
	for _, command := range commands {
		if matchPlatform(command.TargetPlatforms, os, arch) {
			filteredCommands = append(filteredCommands, command)
		}
	}
	return filteredCommands
}

func matchPlatform(platformOptions []string, os string, arch string) bool {
	if len(platformOptions) == 0 {
		return true
	}
	for _, platform := range platformOptions {
		if system.MatchesPlatform(platform, os, arch) {
			return true
		}
	}
	return false
}
