import { MouseEventHandler } from "react";

export default function TrashIcon({ className, onClick }: { className?: string, onClick?: MouseEventHandler<SVGSVGElement> | undefined }) {
    return (
        <svg
            className={className}
            onClick={onClick}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="M3 6h18" /><path d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2" /><path d="M10 11v6" /><path d="M14 11v6" />
        </svg>
    )
}