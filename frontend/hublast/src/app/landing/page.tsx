import Link from "next/link";

export default function Calma() {
    return (
        <div className="min-h-screen flex flex-col justify-center items-center bg-gradient-to-b from-neutral-100 to-white dark:from-gray-950 dark:to-gray-900 px-4 text-center">
            <div className="max-w-md">
                <div className="text-6xl mb-6">🚧</div>
                <h1 className="text-4xl font-bold mb-4 text-gray-800 dark:text-white">
                    Tenha calma
                </h1>
                <h2 className="text-xl text-gray-600 dark:text-gray-300 mb-8">
                    Ainda estamos trabalhando nessa página...
                </h2>
                <Link
                    href="/"
                    className="inline-block bg-violet-600 dark:bg-violet-500 text-white font-medium py-2 px-6 rounded-lg hover:bg-violet-700 dark:hover:bg-violet-600 transition-colors"
                >
                    Voltar para o início
                </Link>
            </div>
        </div>
    );
}
