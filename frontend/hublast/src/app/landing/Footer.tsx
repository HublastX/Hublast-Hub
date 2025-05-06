import { FaGithub, FaLinkedin, FaEnvelope } from "react-icons/fa";

export default function Footer() {
    return (
        <footer className="w-full px-6 py-10 text-center flex flex-col items-center">
            <h3 className="text-lg font-semibold text-purple-600 dark:text-purple-400 mb-4">
                Hublast
            </h3>
            <p className="text-sm mb-4">
                Conectando desenvolvedores, projetos e crescimento contínuo.
            </p>
            <div className="flex gap-6 mb-4">
                <a
                    href="https://github.com/HublastX"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-purple-600 hover:text-purple-800 transition-colors text-xl"
                >
                    <FaGithub />
                </a>
                <a
                    href="https://www.linkedin.com/company/hublast/"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-purple-600 hover:text-purple-800 transition-colors text-xl"
                >
                    <FaLinkedin />
                </a>
                <a
                    href="mailto:hublastx@gmail.com"
                    className="text-purple-600 hover:text-purple-800 transition-colors text-xl"
                >
                    <FaEnvelope />
                </a>
            </div>
            <p className="text-xs text-gray-500 dark:text-gray-500">
                © {new Date().getFullYear()} Hublast. Todos os direitos
                reservados.
            </p>
        </footer>
    );
}
