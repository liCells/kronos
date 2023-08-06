/** @type {import('tailwindcss').Config} */
export default {
    darkMode: 'media',
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        backgroundColor: theme => ({
            ...theme('colors'),
            'dark': '#010409',
            'searchInput': '#0e1116',
        }),
        extend: {},
    },
    plugins: [],
}
