package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kitdevelop-org/vtx/internal/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <plugin-id>",
	Short: "Añade una dependencia a otro plugin de Veritix",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetPluginId := args[0]
		
		cfg, err := config.LoadConfig("vtx.config.json")
		if err != nil {
			return fmt.Errorf("no se encontró vtx.config.json: %w", err)
		}

		fmt.Printf("📦 Añadiendo dependencia a: %s.Contracts...\n", targetPluginId)

		// 1. Añadir el paquete NuGet al proyecto principal
		// El paquete se llama Veritix.Plugin.<PluginId>.Contracts
		packageName := fmt.Sprintf("Veritix.Plugin.%s.Contracts", targetPluginId)
		dotnetCmd := exec.Command("dotnet", "add", cfg.Backend.Project, "package", packageName)
		dotnetCmd.Stdout = os.Stdout
		dotnetCmd.Stderr = os.Stderr
		if err := dotnetCmd.Run(); err != nil {
			return fmt.Errorf("falló al añadir el paquete NuGet: %w", err)
		}

		// 2. Actualizar manifest.json
		if err := updateManifestDependencies(targetPluginId); err != nil {
			return fmt.Errorf("error actualizando manifest.json: %w", err)
		}

		// 3. Actualizar vite.config.ts (si aplica)
		if cfg.HasFrontend {
			if err := updateViteRemotes(cfg.Frontend.Root, targetPluginId); err != nil {
				_ = fmt.Errorf("advertencia: no se pudo actualizar vite.config.ts: %w", err)
			}
		}

		fmt.Printf("✅ Dependencia añadida.\n")
		fmt.Printf("👉 Backend: Inyecta I%sCapability para usar su lógica.\n", targetPluginId)
		if cfg.HasFrontend {
			fmt.Printf("👉 Frontend: Importa desde '%s/Shared' para usar sus componentes.\n", targetPluginId)
		}
		return nil
	},
}

func updateViteRemotes(frontendRoot, pluginId string) error {
	viteConfigPath := filepath.Join(frontendRoot, "vite.config.ts")
	if _, err := os.Stat(viteConfigPath); os.IsNotExist(err) {
		return nil
	}

	content, err := os.ReadFile(viteConfigPath)
	if err != nil {
		return err
	}

	// Buscamos el objeto remotes: { ... }
	pluginIdLower := strings.ToLower(pluginId)
	remoteEntry := fmt.Sprintf("\n        '%s': 'http://localhost:3000/plugins/%s/remoteEntry.js',", pluginIdLower, pluginIdLower)
	
	// Inyección simple por texto para la PoC (en producción usaríamos un parser de AST)
	newContent := strings.Replace(string(content), "remotes: {", "remotes: {"+remoteEntry, 1)

	return os.WriteFile(viteConfigPath, []byte(newContent), 0644)
}

func updateManifestDependencies(pluginId string) error {
	manifestPath := "manifest.json" // El CLI se ejecuta en la raíz del plugin
	
	// Si no existe en la raíz, buscar en la carpeta de backend
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		// Buscamos dinámicamente
		matches, _ := filepath.Glob("Veritix.Plugin.*/manifest.json")
		if len(matches) > 0 {
			manifestPath = matches[0]
		} else {
			return fmt.Errorf("no se encontró manifest.json")
		}
	}

	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return err
	}

	var manifest map[string]interface{}
	if err := json.Unmarshal(data, &manifest); err != nil {
		return err
	}

	// Obtener o crear la lista de dependencias
	deps, ok := manifest["dependencies"].([]interface{})
	if !ok {
		deps = []interface{}{}
	}

	// Evitar duplicados
	for _, d := range deps {
		if d == pluginId {
			return nil // Ya existe
		}
	}

	manifest["dependencies"] = append(deps, pluginId)

	newData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(manifestPath, newData, 0644)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
