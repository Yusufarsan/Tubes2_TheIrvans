import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import path from "path"
import { TanStackRouterVite } from "@tanstack/router-vite-plugin"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), TanStackRouterVite()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
  preview: {
    port: 8000,
    strictPort: true,
  },
  server: {
    port: 8000,
    strictPort: true,
    host: true,
  },
})
