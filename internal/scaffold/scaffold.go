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
	HasContracts   bool
	LicenseType    string
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
		HasContracts:  a.HasContracts,
		LicenseType:   a.LicenseType,
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
		filepath.Join(pluginDir, "Data"),
		filepath.Join(pluginDir, "Licensing"),
	}

	if a.HasContracts {
		dirs = append(dirs, contractsDir)
	}

	if a.HasFrontend {
		dirs = append(dirs, filepath.Join(pluginDir, "ClientApp", "src"))
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creando directorio %s: %w", dir, err)
		}
	}

	// 1. Generar Contracts .csproj (Opcional)
	if a.HasContracts {
		data.ItemGroup = ""
		if err := render(filepath.Join(contractsDir, "Veritix.Plugin."+a.PluginId+".Contracts.csproj"), "templates/backend/csproj.tmpl", data); err != nil {
			return err
		}
	}

	// 2. Generar Main Plugin .csproj
	itemGroup := `
  <ItemGroup>
    <PackageReference Include="Microsoft.EntityFrameworkCore.Design" Version="10.0.7">
      <PrivateAssets>all</PrivateAssets>
      <IncludeAssets>runtime; build; native; contentfiles; analyzers; buildtransitive</IncludeAssets>
    </PackageReference>
    <PackageReference Include="Npgsql.EntityFrameworkCore.PostgreSQL" Version="10.0.0" />
  </ItemGroup>`

	if a.HasContracts {
		itemGroup = fmt.Sprintf(`
  <ItemGroup>
    <ProjectReference Include="..\Veritix.Plugin.%s.Contracts\Veritix.Plugin.%s.Contracts.csproj" />
  </ItemGroup>`, a.PluginId, a.PluginId) + itemGroup
	}

	data.ItemGroup = itemGroup
	if err := render(filepath.Join(pluginDir, "Veritix.Plugin."+a.PluginId+".csproj"), "templates/backend/csproj.tmpl", data); err != nil {
		return err
	}

	// 3. Generar Código C#
	backendFiles := map[string]string{
		filepath.Join(pluginDir, a.PluginId+"Plugin.cs"):            "templates/backend/PluginBase.cs.tmpl",
		filepath.Join(pluginDir, "Controllers", a.PluginId+"Controller.cs"): "templates/backend/Controller.cs.tmpl",
		filepath.Join(pluginDir, "Services", "I"+a.PluginId+"Service.cs"):   "templates/backend/Service.cs.tmpl",
		filepath.Join(pluginDir, "Data", a.PluginId+"DbContext.cs"):         "templates/backend/DbContext.cs.tmpl",
		filepath.Join(pluginDir, "Data", a.PluginId+"Migrator.cs"):          "templates/backend/Migrator.cs.tmpl",
	}

	if a.HasContracts {
		backendFiles[filepath.Join(contractsDir, "I"+a.PluginId+"Capability.cs")] = "templates/backend/Capability.cs.tmpl"
	}

	if a.LicenseType == "Custom" {
		backendFiles[filepath.Join(pluginDir, "Licensing", a.PluginId+"LicenseValidator.cs")] = "templates/backend/LicenseValidator.cs.tmpl"
	}

	for dest, tmpl := range backendFiles {
		if err := render(dest, tmpl, data); err != nil {
			return fmt.Errorf("error generando archivo backend %s: %w", dest, err)
		}
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
		
		// Archivos base del frontend
		files := map[string]string{
			filepath.Join(clientAppDir, "index.html"):          "templates/frontend/index.html.tmpl",
			filepath.Join(clientAppDir, "package.json"):        "templates/frontend/package.json.tmpl",
			filepath.Join(clientAppDir, "vite.config.ts"):      "templates/frontend/vite.config.ts.tmpl",
			filepath.Join(clientAppDir, "tsconfig.json"):       "templates/frontend/tsconfig.json.tmpl",
			filepath.Join(clientAppDir, "tailwind.config.js"):  "templates/frontend/tailwind.config.js.tmpl",
			filepath.Join(clientAppDir, "postcss.config.js"):   "templates/frontend/postcss.config.js.tmpl",
			filepath.Join(clientAppDir, "src", "main.tsx"):      "templates/frontend/main.tsx.tmpl",
			filepath.Join(clientAppDir, "src", "index.css"):     "templates/frontend/index.css.tmpl",
			filepath.Join(clientAppDir, "src", "PluginApp.tsx"): "templates/frontend/PluginApp.tsx.tmpl",
		}

		for dest, tmpl := range files {
			if err := render(dest, tmpl, data); err != nil {
				return fmt.Errorf("error generando archivo frontend %s: %w", dest, err)
			}
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
