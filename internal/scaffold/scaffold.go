package scaffold

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kitdevelop-org/vtx/internal/ui"
)

//go:embed templates/*
var templatesFS embed.FS

type templateData struct {
	PluginId       string
	PluginIdLower  string
	DisplayName    string
	Author         string
	Country        string
	HasFrontend    bool
	BackendProject string
	FrontendRoot   string
	ItemGroup      string
}

func Generate(a *ui.NewPluginAnswers) error {
	baseDir := strings.ToLower(strings.ReplaceAll(a.Name, " ", "-"))
	
	data := templateData{
		PluginId:      a.PluginId,
		PluginIdLower: strings.ToLower(a.PluginId),
		DisplayName:   a.Name,
		Author:        a.Author,
		Country:       a.Country,
		HasFrontend:   a.HasFrontend,
		BackendProject: fmt.Sprintf("./Veritix.Plugin.%s/Veritix.Plugin.%s.csproj", a.PluginId, a.PluginId),
		FrontendRoot:   fmt.Sprintf("./Veritix.Plugin.%s/ClientApp", a.PluginId),
	}

	// Estructura de carpetas
	pluginDir := filepath.Join(baseDir, "Veritix.Plugin."+a.PluginId)
	contractsDir := filepath.Join(baseDir, "Veritix.Plugin."+a.PluginId+".Contracts")
	
	dirs := []string{
		pluginDir,
		filepath.Join(pluginDir, "Controllers"),
		filepath.Join(pluginDir, "Services"),
		contractsDir,
	}

	if a.HasFrontend {
		dirs = append(dirs, filepath.Join(pluginDir, "ClientApp", "src"))
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creando directorio %s: %w", dir, err)
		}
	}

	// 1. Generar Contracts .csproj
	data.ItemGroup = ""
	if err := render(filepath.Join(contractsDir, "Veritix.Plugin."+a.PluginId+".Contracts.csproj"), "templates/backend/csproj.tmpl", data); err != nil {
		return err
	}

	// 2. Generar Main Plugin .csproj
	data.ItemGroup = fmt.Sprintf(`
  <ItemGroup>
    <ProjectReference Include="..\Veritix.Plugin.%s.Contracts\Veritix.Plugin.%s.Contracts.csproj" />
  </ItemGroup>`, a.PluginId, a.PluginId)
	if err := render(filepath.Join(pluginDir, "Veritix.Plugin."+a.PluginId+".csproj"), "templates/backend/csproj.tmpl", data); err != nil {
		return err
	}

	// 3. Generar PluginBase.cs
	if err := render(filepath.Join(pluginDir, a.PluginId+"Plugin.cs"), "templates/backend/PluginBase.cs.tmpl", data); err != nil {
		return err
	}

	// 4. Generar vtx.config.json
	if err := render(filepath.Join(baseDir, "vtx.config.json"), "templates/vtx.config.json.tmpl", data); err != nil {
		return err
	}

	// 5. Generar manifest.json
	if err := render(filepath.Join(pluginDir, "manifest.json"), "templates/manifest.json.tmpl", data); err != nil {
		return err
	}

	// 6. Generar Frontend (si aplica)
	if a.HasFrontend {
		clientAppDir := filepath.Join(pluginDir, "ClientApp")
		if err := render(filepath.Join(clientAppDir, "src", "PluginApp.tsx"), "templates/frontend/PluginApp.tsx.tmpl", data); err != nil {
			return err
		}
		// Nota: Para package.json podríamos usar otra plantilla o mantener el string format
		if err := createFrontendPackageJson(clientAppDir, a); err != nil {
			return err
		}
	}

	return createNugetConfig(baseDir)
}

func render(destPath, tmplPath string, data interface{}) error {
	tmpl, err := template.ParseFS(templatesFS, tmplPath)
	if err != nil {
		return err
	}

	f, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

// Estos dos podrían moverse a plantillas también para ser 100% consistentes
func createNugetConfig(dir string) error {
	path := filepath.Join(dir, "nuget.config")
	content := `<?xml version="1.0" encoding="utf-8"?>
<configuration>
  <packageSources>
    <clear />
    <add key="nuget.org" value="https://api.nuget.org/v3/index.json" />
    <add key="kitdevelop" value="https://nuget.pkg.github.com/kitdevelop-org/index.json" />
  </packageSources>
  <packageSourceCredentials>
    <kitdevelop>
      <add key="Username" value="YOUR_GITHUB_USER" />
      <add key="ClearTextPassword" value="YOUR_PAT_TOKEN" />
    </kitdevelop>
  </packageSourceCredentials>
</configuration>`
	return os.WriteFile(path, []byte(content), 0644)
}

func createFrontendPackageJson(dir string, a *ui.NewPluginAnswers) error {
	path := filepath.Join(dir, "package.json")
	content := fmt.Sprintf(`{
  "name": "@plugins/%s-ui",
  "version": "1.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "preview": "vite preview"
  },
  "dependencies": {
    "react": "^19.0.0",
    "react-dom": "^19.0.0",
    "lucide-react": "^0.450.0",
    "@kitdevelop-org/veritix-ui-kit": "^0.1.0"
  },
  "devDependencies": {
    "@types/react": "^19.0.0",
    "@types/react-dom": "^19.0.0",
    "@vitejs/plugin-react": "^4.3.0",
    "typescript": "^5.6.0",
    "vite": "^6.0.0",
    "tailwindcss": "^3.4.0",
    "autoprefixer": "^10.4.0",
    "postcss": "^8.4.0"
  }
}`, strings.ToLower(a.PluginId))
	return os.WriteFile(path, []byte(content), 0644)
}
