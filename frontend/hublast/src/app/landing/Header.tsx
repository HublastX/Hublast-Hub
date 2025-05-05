import Link from "next/link";
import Logo from "../assets/Logo";
import gsap from "gsap";
import { ScrollToPlugin } from "gsap/ScrollToPlugin";

gsap.registerPlugin(ScrollToPlugin);

export default function Header() {
    const listHeader = [
        { text: "Início", route: "#start" },
        { text: "Sobre", route: "#about" },
    ];

    const handleClick = (
        e: React.MouseEvent<HTMLAnchorElement>,
        route: string
    ) => {
        e.preventDefault();
        const id = route.replace("#", "");
        const target = document.getElementById(id);
        if (target) {
            const header = document.querySelector("header");
            const headerOffset = header?.clientHeight || 80;
            const elementPosition =
                target.getBoundingClientRect().top + window.pageYOffset;
            const offsetPosition = elementPosition - headerOffset;

            gsap.to(window, {
                duration: 1,
                scrollTo: {
                    y: offsetPosition,
                    autoKill: true,
                },
                ease: "power2.out",
            });
        }
    };

    return (
        <header className="fixed z-50 w-full py-4 px-6 flex justify-between items-center bg-white/70 dark:bg-gray-900/70 backdrop-blur-md border-b border-white/10 dark:border-none">
            <div className="flex items-center">
                <Logo className="h-12 text-purple-500" />
                <h1 className="text-2xl rounded px-3 py-2 font-extrabold">
                    Hublast
                </h1>
            </div>

            <nav className="flex gap-6 px-3 py-2 rounded md:h-fit items-center">
                {listHeader.map((item, index) => (
                    <Link
                        key={index}
                        href={item.route}
                        onClick={(e) => handleClick(e, item.route)}
                        className="hover:text-purple-500 transition-colors duration-200 font-medium"
                    >
                        {item.text}
                    </Link>
                ))}
            </nav>
        </header>
    );
}
