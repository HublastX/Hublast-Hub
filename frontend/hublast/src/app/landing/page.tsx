import Head from "next/head";
import Header from "./Header";
import About from "./about";
import { FaDiscord } from "react-icons/fa";
import Footer from "./Footer";

export default function Landing() {
    return (
        <>
            <Head>
                <title>Hublast | Comunidade Dev</title>
            </Head>

            <main className="flex flex-col items-center min-h-screen scroll-smooth">
                <Header />
                <section id="start" className="w-full h-screen flex flex-col justify-center items-center text-center px-4 bg-gradient-to-b from-transparent to-neutral-50 dark:to-gray-950 scroll-smooth">
                    <p className="uppercase tracking-widest text-sm md:text-base text-gray-400 mb-4">
                        a comunidade que acredita no seu potencial
                    </p>

                    <h1 className="text-4xl md:text-6xl font-extrabold leading-tight mb-4">
                        Sua <span className="text-purple-500">jornada dev</span>{" "}
                        <br />
                        começa com{" "}
                        <span className="text-purple-400">apoio e prática</span>
                    </h1>

                    <p className="text-md md:text-lg text-gray-300 mb-8 max-w-xl">
                        Aprenda, colabore e ganhe experiência com projetos
                        reais.
                    </p>

                    <a
                        href="https://discord.gg/seulink"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="inline-flex items-center gap-2 bg-purple-600 hover:bg-purple-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-300"
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
