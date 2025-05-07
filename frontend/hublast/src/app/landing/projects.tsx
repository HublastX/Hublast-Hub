"use client";
import { useEffect, useRef } from "react";
import { gsap } from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";
import { FaExternalLinkAlt } from "react-icons/fa";

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
    ];

    useEffect(() => {
        gsap.fromTo(
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
    }, []);

    return (
        <section
            id="projects"
            className="w-full px-6 py-20 bg-gradient-to-b from-neutral-50 dark:from-gray-950 to-transparent text-left flex flex-col items-center scroll-smooth"
        >
            <h2 className="text-3xl md:text-4xl font-semibold mb-16 text-gray-900 dark:text-white text-center">
                Projetos
            </h2>

            <div className="flex flex-col gap-10 max-w-4xl w-full grid-cols-2">
                {projects.map((project, index) => (
                    <div
                        key={index}
                        ref={(el) => {
                            if (el) projectsRef.current[index] = el;
                        }}
                        className="relative bg-white dark:bg-gray-900 border-l-4 border-purple-500 shadow-md p-6 rounded-lg opacity-0 hover:shadow-xl transition-shadow duration-300"
                    >
                        <h3 className="text-2xl font-bold text-purple-600 mb-2">
                            {project.title}
                        </h3>
                        <p className="text-gray-600 dark:text-gray-300 mb-4">
                            {project.description}
                        </p>
                        <a
                            href={project.link}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="inline-flex items-center text-purple-600 hover:underline font-medium"
                        >
                            Ver no GitHub{" "}
                            <FaExternalLinkAlt className="ml-2" size={14} />
                        </a>
                    </div>
                ))}
            </div>
        </section>
    );
}
