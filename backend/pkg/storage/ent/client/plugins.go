package client

import (
	"context"

	v1alphaPlugins "github.com/redhat-appstudio/quality-studio/api/apis/plugins/v1alpha1"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/plugins"
)

// CreatePlugin populate database with a new given plugin.
func (d *Database) CreatePlugin(plugin *v1alphaPlugins.PluginSpec) (*db.Plugins, error) {
	pluginCreated, err := d.client.Plugins.Create().
		SetName(plugin.Name).
		SetCategory(plugin.Category).
		SetLogo(plugin.Logo).
		SetDescription(plugin.Description).
		SetReason(plugin.Reason).
		Save(context.Background())
	if err != nil {
		return nil, convertDBError("create plugins status: %w", err)
	}

	return pluginCreated, nil
}

// ListPlugins extracts an array with all plugins from the database.
func (d *Database) ListPlugins() ([]*db.Plugins, error) {
	plugins, err := d.client.Plugins.Query().All(context.TODO())
	if err != nil {
		return nil, convertDBError("list plugins: %w", err)
	}

	return plugins, nil
}

// GetPluginByName extract a plugin from database from given name.
func (d *Database) GetPluginByName(pluginName string) (*db.Plugins, error) {
	plugin, err := d.client.Plugins.Query().Where(plugins.Name(pluginName)).Only(context.Background())
	if err != nil {
		return nil, convertDBError("get plugin: %w", err)
	}

	return plugin, nil
}

// GetPluginsByTeam extract all plugins which belong to a given name in database.
func (d *Database) GetPluginsByTeam(team *db.Teams) (storagePlugin []*v1alphaPlugins.Plugin, err error) {
	allPlugins, err := d.client.Plugins.Query().All(context.TODO())
	if err != nil {
		return nil, convertDBError("list plugins: %w", err)
	}

	pluginTeam, err := d.client.Teams.QueryPlugins(team).All(context.Background())

	if err != nil {
		return nil, convertDBError("get plugin: %w", err)
	}

	for _, plugin := range allPlugins {
		storagePlugin = append(storagePlugin, installedPlugin(plugin, pluginInstalled(plugin, pluginTeam)))
	}

	return storagePlugin, nil
}

func pluginInstalled(plugin *db.Plugins, installedPlugins []*db.Plugins) bool {
	for _, pg := range installedPlugins {
		if pg.Name == plugin.Name {
			return true
		}
	}
	return false
}

// InstallPlugin assign a plugin to a team in database
func (d *Database) InstallPlugin(team *db.Teams, plugin *db.Plugins) (db *db.Teams, err error) {
	return d.client.Teams.UpdateOne(team).AddPlugins(plugin).Save(context.Background())
}

// RemovePlugin means that an assigned plugin to a team will be removed
func (d *Database) RemovePlugin(team *db.Teams, plugin *db.Plugins) (*db.Teams, error) {
	return d.client.Teams.UpdateOne(team).RemovePlugins(plugin).Save(context.Background())
}