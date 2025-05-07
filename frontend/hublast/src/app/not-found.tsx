"use client";
import Link from "next/link";
import { FaArrowLeft } from "react-icons/fa";

export default function NotFound() {
    return (
        <section className="w-full min-h-screen flex flex-col justify-center items-center px-6 py-20 bg-neutral-100 dark:bg-gray-950 text-center">
            <h1 className="text-7xl font-extrabold text-violet-600 mb-4">
                404
            </h1>
            <h2 className="text-2xl md:text-3xl font-semibold text-gray-800 dark:text-white mb-4">
                Página não encontrada
            </h2>
            <p className="text-gray-600 dark:text-gray-400 mb-8 max-w-md">
                A página que você está tentando acessar não existe ou foi
                movida.
            </p>
            <Link
                href="/"
                className="inline-flex items-center gap-2 bg-violet-600 hover:bg-violet-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-200"
            >
                <FaArrowLeft />
                Voltar para o início
            </Link>
        </section>
    );
}
