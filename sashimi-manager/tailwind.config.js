/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        cabin: 'cabin, serif'
      },
      colors: {
        'sashimi-blue': '#E3F5FF',
        'sashimi-gray': '#E5ECF6',
        'sashimi-pink': '#ffe6d3',
        'sashimi-purple': '#95a3fc',
        'sashimi-green': '#e2f4ef',
        'sashimi-deepblue': '#1f77b4',
        'sashimi-deeppink': '#ff7f0e',
        'sashimi-deepgreen': '#006400',
        'sashimi-deeppurple': '#4B0082',
        'sashimi-deepgray': '#505050'
      }
    }
  },
  plugins: []
};
