/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["*.{html,js}", "./**/*.{html,js}", "../templates/**/*.{html,js}"],
    theme: {
        extend: {
            fontFamily: {
                pigarnosneue: [
                    '"PIGARNOS NEUE Regular"',
                    {
                        fontFeatureSettings: '"cv11", "ss01"',
                        fontVariationSettings: '"opsz" 32',
                    },
                ],
                jasminelaslo: [
                    '"Jasmine Laslo Regular"',
                    {
                        fontFeatureSettings: '"cv11", "ss01"',
                        fontVariationSettings: '"opsz" 32',
                    },
                ],
                brightania: [
                    '"Brightania"',
                    {
                        fontFeatureSettings: '"cv11", "ss01"',
                        fontVariationSettings: '"opsz" 32',
                    },
                ],
                hightnight: [
                    '"Hight Night Regular"',
                    {
                        fontFeatureSettings: '"cv11", "ss01"',
                        fontVariationSettings: '"opsz" 32',
                    },
                ],
                helvetica: [
                    '"Helvetica"',
                    {
                        fontFeatureSettings: '"cv11", "ss01"',
                        fontVariationSettings: '"opsz" 32',
                    },
                ],
            },
            colors: {
                transparent: "transparent",
                current: "currentColor",
                subwaylane: {
                    1: "#0055a0",
                    2: "#037f61",
                    3: "#f03f35",
                    4: "#ffd600",
                    5: "#774c9e",
                    6: "#000000",
                    7: "#a61366",
                    8: "#a1a197",
                    9: "#02a78b",
                    10: "#007b94",
                    11: "#ef4c23",
                    12: "#024381",
                    13: "#09ad61",
                    14: "#000000",
                    15: "#848d91",
                },
                subwaylanetext: {
                    1: "#eeeeee",
                    2: "#eeeeee",
                    3: "#eeeeee",
                    4: "#111111",
                    5: "#eeeeee",
                    6: "#000000",
                    7: "#eeeeee",
                    8: "#eeeeee",
                    9: "#eeeeee",
                    10: "#eeeeee",
                    11: "#eeeeee",
                    12: "#eeeeee",
                    13: "#eeeeee",
                    14: "#eeeeee",
                    15: "#eeeeee",
                },
            },
            keyframes: {
                "fade-in": {
                    "0%": {
                        opacity: 0,
                    },
                    "100%": {
                        opacity: 1,
                    },
                },
                "fade-out": {
                    "0%": {
                        opacity: 1,
                    },
                    "100%": {
                        opacity: 0,
                        height: 0,
                    },
                },
            },
            animation: {
                fadein: "fade-in 1s ease-in-out 0.25s 1",
                fadeout: "fade-out 1s ease-out 1s 1 forwards",
            },
        },
        plugins: [],
    },
};
