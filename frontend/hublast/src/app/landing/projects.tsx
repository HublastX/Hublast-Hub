"use client";
import { useEffect, useRef } from "react";
import { gsap } from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";
import { FaExternalLinkAlt, FaGithub } from "react-icons/fa";

gsap.registerPlugin(ScrollTrigger);

export default function Projects() {
    const projectsRef = useRef<HTMLDivElement[]>([]);

    const projects = [
        {
            title: "Commit-IA",
            description:
                "Commit-IA é uma ferramenta de linha de comando criada em Go que utiliza Large Language Models (LLMs) para analisar suas alterações no código Git e gerar automaticamente mensagens de confirmação semântica.",
            link: "https://github.com/HublastX/Commit-IA",
        },
        {
            title: "Book-Guardian",
            description:
                "Projeto BookGuardian criado com a função de gerencia livros , que visa auxiliar os amantes da leitura a gerenciar suas bibliotecas pessoais de maneira eficiente e organizada.",
            link: "https://github.com/HublastX/Book-Guardian",
            depoly: "https://book-guardian-production.up.railway.app/",
        },
        {
            title: "Validgen",
            description:
                "Valigen é uma biblioteca Go para gerar e validar diversos tipos de dados estruturados, como CPF, CNPJ e outros identificadores formatados. Ela fornece uma API eficiente e fácil de usar para lidar com validação, formatação e geração de dados.",
            link: "https://github.com/HublastX/ValidGen",
        },
    ];

    useEffect(() => {
        const animation = gsap.fromTo(
            projectsRef.current,
            { opacity: 0, x: -50 },
            {
                opacity: 1,
                x: 0,
                duration: 1,
                ease: "power3.out",
                stagger: 0.2,
                scrollTrigger: {
                    trigger: "#projects",
                    start: "top 80%",
                },
            }
        );

        return () => {
            ScrollTrigger.getAll().forEach(trigger => trigger.kill());
            animation.kill();
        };
    }, []);

    projectsRef.current = [];

    return (
        <section
            id="projects"
            className="w-full px-6 py-20 bg-gradient-to-b from-neutral-50 dark:from-gray-950 to-transparent text-left flex flex-col items-center scroll-smooth"
        >
            <h2 className="text-3xl md:text-4xl font-semibold mb-16 text-gray-900 dark:text-white text-center">
                Últimos{" "}
                <span className="bg-gradient-to-r from-violet-500 to-indigo-500 bg-clip-text text-transparent font-bold">
                    Projetos
                </span>
            </h2>

            <div className="flex flex-col gap-10 max-w-4xl w-full">
                {projects.map((project, index) => (
                    <div
                        key={index}
                        ref={(el) => {
                            if (el) projectsRef.current[index] = el;
                        }}
                        className="relative bg-white dark:bg-gray-900 border-l-4 border-violet-500 shadow-md p-6 rounded-lg opacity-0 hover:shadow-xl transition-shadow duration-300"
                    >
                        <h3 className="text-2xl font-bold dark:text-violet-500 text-violet-600 mb-2">
                            {project.title}
                        </h3>
                        <p className="text-gray-600 dark:text-gray-300 mb-4">
                            {project.description}
                        </p>
                        <div className="flex flex-wrap gap-6 w-full">
                            <a
                                href={project.link}
                                target="_blank"
                                rel="noopener noreferrer"
                                className="inline-flex items-center dark:text-violet-500 text-violet-600 hover:underline font-medium"
                            >
                                <FaGithub className="mr-2" size={20} /> Ver no GitHub
                            </a>
                            {project.depoly && (
                                <a
                                    href={project.depoly}
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    className="inline-flex items-center dark:text-violet-500 text-violet-600 hover:underline font-medium"
                                >
                                    Visitar o site <FaExternalLinkAlt className="ml-2" size={14} />
                                </a>
                            )}
                        </div>
                    </div>
                ))}
            </div>

            <a
                href="/landing"
                className="mt-12 text-violet-600 dark:text-violet-500 hover:underline font-medium text-lg"
            >
                Veja todos os nossos projetos
            </a>
        </section>
    );
}
