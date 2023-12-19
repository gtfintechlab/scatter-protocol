"use client";

import { useEffect, useState, useContext } from "react";
import BoxIcon from "@/components/icons/BoxIcon";
import Link from "next/link";
import { MouseEvent } from "react";
import SandwichMenuIcon from "@/components/icons/SandwichMenuIcon";
import { NAVBAR_ITEMS } from "@/utils/constants";
import { ScreenContext } from "@/contexts/ScreenContext";
import { NavbarItem } from "@/utils/types";

export default function NavigationBar() {
    const [toggleOpen, setToggleOpen] = useState<boolean>(false);
    const [initialLoad, setInitialLoad] = useState<boolean>(false);
    const { isMobile, currentScreen } = useContext(ScreenContext);

    const toggleNavbar = (event: MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        setToggleOpen(!toggleOpen);
    };

    useEffect(() => {
        const handleResize = () => {
            // 768 pixels for sandwich header to show up
            const mobileWidthThreshold = 768;
            if (window.innerWidth > mobileWidthThreshold) {
                setToggleOpen(false);
            }
        };
        handleResize();
        window.addEventListener('resize', handleResize);
        setInitialLoad(true);
        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, []);

    return (
        <div className="relative">
            <header className="flex items-center h-16 px-4 border-b bg-white shadow-sm text-gray-600 md:px-6">
                <nav className="flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 lg:gap-6">
                    {!isMobile && <Link className="flex items-center gap-2 text-lg font-semibold md:text-base"
                        href="/">
                        <BoxIcon className="h-6 w-6 text-black" />
                        <span className="sr-only">scatter protocol</span>
                    </Link>
                    }
                    {
                        !isMobile && (NAVBAR_ITEMS).map((item, index) => (
                            <Link className={currentScreen == item.url ? "text-gray-600" : "text-gray-400"}
                                href={"/" + item.url} key={index}>
                                {item.text}
                            </Link>
                        ))
                    }
                    {isMobile && initialLoad && <div className="pt-4 pb-4 pr-4">
                        <button onClick={toggleNavbar} className="hover:bg-gray-100 rounded-sm p-0.5">
                            <SandwichMenuIcon className="w-6 h-6"></SandwichMenuIcon>
                        </button>
                    </div>
                    }
                </nav>
            </header>
            {
                isMobile && toggleOpen && <div className="flex flex-col absolute w-full z-50">
                    {
                        (NAVBAR_ITEMS).map((item: NavbarItem, index: number) => (
                            <Link className={`pl-4 pt-2 pb-2 bg-white font-medium ${(currentScreen == item.url ? "text-gray-600" : "text-gray-400")}`}
                                href={"/" + item.url} key={index}>
                                {item.text}
                            </Link>
                        ))}
                </div>
            }
        </div>
    )
}
