// vite.config.js
import tailwindcss from "tailwindcss";

export default {
  // config options
  base: "/",
  css: {
    postcss: {
      plugins: [tailwindcss()],
    },
  },
};
