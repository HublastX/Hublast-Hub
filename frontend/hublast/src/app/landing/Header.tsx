import Image from "next/image";
import Link from "next/link";

export default function Header() {
  const listHeader = [
    { text: "Início", route: "#start" },
    { text: "Sobre", route: "#about" },
    ];

  return (
    <header className="fixed w-full py-4 px-6 flex justify-between items-cente backdrop-blur-md md:backdrop-blur-none">
      <div className="flex items-center gap-4">
        <Image
          src="https://avatars.githubusercontent.com/u/155011581?s=200&v=4"
          alt="Ícone da Hublast"
          width={60}
          height={60}
          className="rounded-full"
        />
        <h1 className="text-2xl rounded pr-3 py-2 font-extrabold md:backdrop-blur-md">Hublast</h1>
      </div>

      <nav className="flex gap-6 pl-3 py-2 rounded md:backdrop-blur-md md:h-fit items-center">
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
