interface LogoProps {
    className?: string;
}

export default function Logo({ className }: LogoProps) {
    return (
        <svg
            className={className}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="129.612 24.354 452.353 690.829"
        >
            <ellipse
                style={{
                    stroke: "currentColor",
                    strokeWidth: 30,
                    fill: "none",
                }}
                cx="353.616"
                cy="374.663"
                rx="200"
                ry="200"
            ></ellipse>

            <path
                style={{
                    stroke: "currentColor",
                    fill: "none",
                    strokeLinecap: "round",
                    strokeLinejoin: "round",
                    strokeWidth: 20,
                }}
                d="M 343.381 166.334 C 342.759 159.565 345.008 61.63 345.393 61.47 C 365.562 53.133 393.666 171.534 365.078 357.11 C 336.49 542.686 377.951 695.226 383.709 695.348 C 388.823 695.457 391.297 592.898 390.561 578.536"
            ></path>
        </svg>
    );
}

