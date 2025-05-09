"use client";
import { useState } from "react";
import Link from "next/link";
import Logo from "../../assets/Logo";
import gsap from "gsap";
import { ScrollToPlugin } from "gsap/ScrollToPlugin";
import { HiMenu, HiX } from "react-icons/hi";

gsap.registerPlugin(ScrollToPlugin);

export default function Header() {
    const [menuOpen, setMenuOpen] = useState(false);

    const links = [
        { text: "Início", route: "#start" },
        { text: "Sobre", route: "#about" },
        { text: "Projetos", route: "#projects" },
    ];

    const handleClick = (
        e: React.MouseEvent<HTMLAnchorElement>,
        route: string
    ) => {
        e.preventDefault();
        setMenuOpen(false);

        const id = route.replace("#", "");
        const target = document.getElementById(id);

        if (target) {
            const headerHeight = document.querySelector("header")?.clientHeight || 80;
            const position = target.getBoundingClientRect().top + window.pageYOffset - headerHeight;

            gsap.to(window, {
                duration: 1,
                scrollTo: { y: position, autoKill: true },
                ease: "power2.out",
            });
        }
    };

    return (
        <header className="fixed z-50 w-full py-4 px-6 flex justify-between items-center bg-white/70 dark:bg-gray-900/70 backdrop-blur-md border-b border-white/10 dark:border-none">
            <div className="flex items-center">
                <Logo className="h-12 text-violet-500" />
                <h1 className="text-2xl font-extrabold px-3 py-2 rounded">
                    Hublast
                </h1>
            </div>

            <button
                className="md:hidden text-3xl"
                onClick={() => setMenuOpen(!menuOpen)}
                aria-label={menuOpen ? "Fechar menu" : "Abrir menu"}
            >
                {menuOpen ? <HiX /> : <HiMenu />}
            </button>

            <nav className="hidden md:flex gap-6 items-center px-3 py-2 rounded">
                {links.map(({ text, route }, index) => (
                    <Link
                        key={index}
                        href={route}
                        onClick={(e) => handleClick(e, route)}
                        className="hover:text-violet-500 transition-colors duration-200 font-medium"
                    >
                        {text}
                    </Link>
                ))}
            </nav>

            {menuOpen && (
                <div className="absolute top-20 right-6 md:hidden flex flex-col gap-4 px-6 py-4 rounded-lg shadow-lg bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700">
                    {links.map(({ text, route }, index) => (
                        <Link
                            key={index}
                            href={route}
                            onClick={(e) => handleClick(e, route)}
                            className="text-gray-700 dark:text-gray-200 hover:text-violet-600 transition-colors duration-200 font-medium"
                        >
                            {text}
                        </Link>
                    ))}
                </div>
            )}
        </header>
    );
}
