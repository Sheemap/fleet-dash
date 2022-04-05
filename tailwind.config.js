module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      gridTemplateRows: {
        'dash': 'repeat(3, 160px)',
      }
    },
  },
  plugins: [],
}
