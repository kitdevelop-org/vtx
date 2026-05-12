import React from 'react';
import { Button, Card } from '@kitdevelop-org/veritix-ui-kit';
import { Box, Globe } from 'lucide-react';

export default function PluginApp() {
  return (
    <div className="p-6 space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-white flex items-center gap-2">
            <Box className="text-veritix-500" />
            Frontend Pro Plugin
          </h1>
          <p className="text-gray-400">Bienvenido al módulo oficial de Frontend Pro.</p>
        </div>
        <Button variant="primary">
          Sincronizar Datos
        </Button>
      </div>
      
      <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
        <Card title="Estado Fiscal (MX)">
          <div className="flex items-center gap-3">
            <Globe className="text-veritix-400" />
            <span>Módulo configurado y operativo.</span>
          </div>
        </Card>

        <Card title="Resumen de Operaciones">
          <p className="text-sm">No hay actividad reciente en este plugin.</p>
        </Card>
      </div>
    </div>
  );
}
