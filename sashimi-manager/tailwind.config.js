/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        cabin: 'cabin, Times New Roman, serif'
      }
    }
  },
  plugins: []
};
