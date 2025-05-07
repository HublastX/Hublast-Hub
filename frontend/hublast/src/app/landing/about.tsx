import { useEffect, useRef } from "react";
import { gsap } from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";
import { FaUsers, FaProjectDiagram, FaLightbulb } from "react-icons/fa";

gsap.registerPlugin(ScrollTrigger);

export default function About() {
    const cardsRef = useRef<HTMLDivElement[]>([]);

    const cards = [
        {
            icon: <FaUsers size={32} className="text-violet-600" />,
            title: "Comunidade ativa",
            description:
                "Conecte-se com devs iniciantes e experientes, troque ideias e evolua em grupo.",
        },
        {
            icon: <FaProjectDiagram size={32} className="text-violet-600" />,
            title: "Projetos reais",
            description:
                "Participe de squads com foco em prática e ganhe experiência colaborando em projetos.",
        },
        {
            icon: <FaLightbulb size={32} className="text-violet-600" />,
            title: "Mentoria e crescimento",
            description:
                "Receba orientações, tire dúvidas e estruture sua carreira com apoio da comunidade.",
        },
    ];

    useEffect(() => {
        gsap.fromTo(
            cardsRef.current,
            { opacity: 0, y: 50 },
            {
                opacity: 1,
                y: 0,
                duration: 1,
                ease: "power3.out",
                stagger: 0.2,
                scrollTrigger: {
                    trigger: "#about",
                    start: "top 80%",
                },
            }
        );
    }, []);

    return (
        <section id="about" className="w-full z-0 px-6 py-20 dark:bg-gray-950 bg-neutral-50 text-center flex flex-col items-center scroll-smooth ">
            <h2 className="text-3xl md:text-4xl font-semibold mb-12 text-gray-900 dark:text-white">
                Sobre a <span className="text-violet-600">Hublast</span>
            </h2>

            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-8 max-w-6xl w-full">
                {cards.map((card, index) => (
                    <div
                        key={index}
                        ref={(el) => {
                            if (el) cardsRef.current[index] = el;
                          }}
                                                  className="bg-white dark:bg-gray-900 p-6 rounded-2xl shadow-lg hover:shadow-xl transition-shadow  duration-300 opacity-0"
                    >
                        <div className="mb-6 flex justify-center items-center">
                            {card.icon}
                        </div>
                        <h3 className="text-xl font-semibold mb-4 text-violet-500">
                            {card.title}
                        </h3>
                        <p className="dark:text-gray-300 text-gray-600 text-sm">
                            {card.description}
                        </p>
                    </div>
                ))}
            </div>
        </section>
    );
}
