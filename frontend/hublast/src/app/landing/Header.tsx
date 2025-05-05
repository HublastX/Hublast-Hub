import Link from "next/link";
import Logo from "../styles/Logo";

export default function Header() {
  const listHeader = [
    { text: "Início", route: "#start" },
    { text: "Sobre", route: "#about" },
    ];

  return (
    <header className="fixed w-full py-4 px-6 flex justify-between items-cente backdrop-blur-md md:backdrop-blur-none">
      <div className="flex items-center">
        <Logo className="h-15 w-15 text-purple-500"  />
        <h1 className="text-2xl rounded px-3 py-2 font-extrabold md:backdrop-blur-md">Hublast</h1>
      </div>

      <nav className="flex gap-6 px-3 py-2 rounded md:backdrop-blur-md md:h-fit items-center">
        {listHeader.map((item, index) => (
          <Link
            key={index}
            href={item.route}
            className="hover:text-purple-500 transition-colors duration-200 font-medium"
          >
            {item.text}
          </Link>
        ))}
      </nav>
    </header>
  );
}
