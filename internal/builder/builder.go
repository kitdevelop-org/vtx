package builder

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kitdevelop-org/vtx/internal/config"
)

func Build(cfg *config.VtxConfig) error {
	// 1. Limpiar y preparar carpetas
	distDir := "./dist"
	publishDir := filepath.Join(distDir, "publish")
	os.RemoveAll(distDir)
	os.MkdirAll(publishDir, 0755)

	// 2. Compilar Backend (.NET 10)
	fmt.Println("📦 Compilando backend (.NET 10)...")
	dotnetCmd := exec.Command("dotnet", "publish", cfg.Backend.Project, "-c", "Release", "-o", publishDir)
	dotnetCmd.Stdout = os.Stdout
	dotnetCmd.Stderr = os.Stderr
	if err := dotnetCmd.Run(); err != nil {
		return fmt.Errorf("falló la compilación de .NET: %w", err)
	}

	// 3. Compilar Frontend (si aplica)
	frontendDist := ""
	if cfg.HasFrontend {
		fmt.Println("🎨 Compilando frontend (pnpm)...")
		pnpmCmd := exec.Command("pnpm", "--dir", cfg.Frontend.Root, "build")
		pnpmCmd.Stdout = os.Stdout
		pnpmCmd.Stderr = os.Stderr
		if err := pnpmCmd.Run(); err != nil {
			return fmt.Errorf("falló la compilación de frontend: %w", err)
		}
		frontendDist = filepath.Join(cfg.Frontend.Root, "dist")
	}

	// 4. Crear el archivo .vtx (ZIP)
	vtxFileName := fmt.Sprintf("%s-%s.vtx", cfg.PluginId, cfg.Version)
	vtxFilePath := filepath.Join(distDir, vtxFileName)
	fmt.Printf("🗜️  Generando paquete: %s...\n", vtxFileName)

	vtxFile, err := os.Create(vtxFilePath)
	if err != nil {
		return err
	}
	defer vtxFile.Close()

	archive := zip.NewWriter(vtxFile)
	defer archive.Close()

	// 5. Agregar manifest.json al root del ZIP
	if err := addFileToZip(archive, "manifest.json", "manifest.json"); err != nil {
		return fmt.Errorf("error agregando manifest.json: %w", err)
	}

	// 6. Agregar binarios de .NET al ZIP (En la raíz, no en bin/)
	if err := addDirToZip(archive, publishDir, ""); err != nil {
		return fmt.Errorf("error agregando binarios: %w", err)
	}

	// 7. Agregar assets de frontend al ZIP
	if cfg.HasFrontend {
		if err := addDirToZip(archive, frontendDist, "wwwroot"); err != nil {
			return fmt.Errorf("error agregando frontend assets: %w", err)
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string, zipPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer, err := zipWriter.Create(zipPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

func addDirToZip(zipWriter *zip.Writer, srcPath string, targetPath string) error {
	return filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		relPath, _ := filepath.Rel(srcPath, path)
		zipPath := filepath.Join(targetPath, relPath)

		return addFileToZip(zipWriter, path, zipPath)
	})
}
