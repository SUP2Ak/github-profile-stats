import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    port: 4173,
    host: true,
    proxy: {
      '/api': {
        target: 'https://github-profile-stats.onrender.com',
        changeOrigin: true,
        secure: false,
      },
    },
  },
});