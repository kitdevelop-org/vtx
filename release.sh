#!/bin/bash
# Script de compilación Multi-OS y Multi-Arquitectura para vtx CLI

APP_NAME="vtx"
OUTPUT_DIR="releases"

# Limpiar versiones anteriores
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Lista de arquitecturas objetivo (OS/ARCH/VARIANT)
# El variant es opcional (ej: para ARMv7)
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "linux/arm/7"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/arm64"
)

echo "🚀 Iniciando compilación global para Veritix CLI (incluyendo soporte ARM/Raspberry Pi)..."

for PLATFORM in "${PLATFORMS[@]}"
do
    OS=$(echo $PLATFORM | cut -d'/' -f1)
    ARCH=$(echo $PLATFORM | cut -d'/' -f2)
    VARIANT=$(echo $PLATFORM | cut -d'/' -f3)
    
    BINARY_NAME="${APP_NAME}-${OS}-${ARCH}"
    if [ ! -z "$VARIANT" ]; then
        BINARY_NAME="${BINARY_NAME}v${VARIANT}"
    fi
    
    if [ "$OS" == "windows" ]; then
        BINARY_NAME="${BINARY_NAME}.exe"
    fi
    
    echo "📦 Compilando para $OS ($ARCH ${VARIANT:-})..."
    
    # Configuración de variables de entorno para Go
    export GOOS=$OS
    export GOARCH=$ARCH
    if [ "$ARCH" == "arm" ] && [ "$VARIANT" == "7" ]; then
        export GOARM=7
    else
        unset GOARM
    fi
    
    go build -o "${OUTPUT_DIR}/${BINARY_NAME}" main.go
done

echo "✅ Compilación completada. Archivos disponibles en la carpeta /releases"
ls -lh $OUTPUT_DIR
