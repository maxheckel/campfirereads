/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {

    extend: {
      colors: {
        'np-yellow': '#D7C17F',
        'np-green': '#4B644A',
        'np-yellow-200': '#F0E9CF',
        'np-dark-brown': '#221B1C',
      },
    },
  },
  plugins: [],
}
