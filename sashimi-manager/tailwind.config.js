/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        cabin: 'cabin, Times New Roman, serif'
      },
      colors: {
        'sashimi-blue': '#E3F5FF',
        'sashimi-gray': '#E5ECF6'
      }
    }
  },
  plugins: []
};
