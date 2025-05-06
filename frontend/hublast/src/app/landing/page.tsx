"use client";
import Head from "next/head";
import Header from "./Header";
import About from "./about";
import { FaDiscord } from "react-icons/fa";
import Footer from "./Footer";
import { useEffect, useRef } from "react";
import { gsap } from "gsap";

export default function Landing() {
    const titleRef = useRef(null);
    const subtitleRef = useRef(null);
    const subtitleDescRef = useRef(null);
    const buttonRef = useRef(null);

    useEffect(() => {
        const tl = gsap.timeline({
            defaults: { ease: "power3.out", duration: 1 },
        });

        tl.fromTo(titleRef.current, { opacity: 0, y: 50 }, { opacity: 1, y: 0 })
            .fromTo(
                subtitleRef.current,
                { opacity: 0, y: 50 },
                { opacity: 1, y: 0 },
                "-=0.8"
            )
            .fromTo(
                subtitleDescRef.current,
                { opacity: 0, y: 50 },
                { opacity: 1, y: 0 },
                "-=9"
            )
            .fromTo(
                buttonRef.current,
                { opacity: 0, y: 50 },
                { opacity: 1, y: 0 },
                "-=0.9"
            );
    }, []);

    return (
        <>
            <Head>
                <title>Hublast | Comunidade Dev</title>
            </Head>

            <main className="flex flex-col items-center min-h-screen scroll-smooth">
                <Header />
                <section
                    id="start"
                    className="w-full h-screen flex flex-col justify-center items-center text-center md:px-4 px-5 bg-gradient-to-b from-transparent to-neutral-50 dark:to-gray-950 scroll-smooth"
                >
                    <p
                        ref={subtitleRef}
                        className="opacity-0 uppercase tracking-widest text-sm md:text-base dark:text-gray-400 text-gray-500 mb-4"
                    >
                        a comunidade que acredita no seu potencial
                    </p>

                    <h1
                        ref={titleRef}
                        className=" opacity-0 text-4xl md:text-6xl font-extrabold leading-tight mb-4"
                    >
                        Sua <span className="text-purple-500">jornada dev</span>{" "}
                        <br />
                        começa com{" "}
                        <span className="text-purple-600">apoio e prática</span>
                    </h1>

                    <p
                        ref={subtitleDescRef}
                        className="opacity-0 text-md md:text-lg dark:text-gray-300 text-gray-400 mb-8 max-w-xl"
                    >
                        Aprenda, colabore e ganhe experiência com projetos
                        reais.
                    </p>

                    <a
                        ref={buttonRef}
                        href="https://discord.gg/uXPXZdkqkf"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="opacity-0 z-0 inline-flex items-center gap-2 bg-purple-600 hover:bg-purple-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-300"
                    >
                        <FaDiscord size={20} />
                        Entrar na Comunidade
                    </a>
                </section>
                <About />
            </main>
            <Footer />
        </>
    );
}
