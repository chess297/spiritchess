import { Footer } from "@/components/Footer";
import { Header } from "@/components/Header";
import type { Metadata } from "next";
import localFont from "next/font/local";
import "@/globals.css";

const geistSans = localFont({
  src: "../fonts/GeistVF.woff",
  variable: "--font-geist-sans",
});
const geistMono = localFont({
  src: "../fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
});

export const metadata: Metadata = {
  title: "广州灵韵星奕软件开发工作室",
  description: "广州灵韵星奕软件开发工作室",
};

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html>
      <body className={`${geistSans} ${geistMono} `}>
        <Header />
        <main className={`flex-auto`}>{children}</main>
        <Footer />
      </body>
    </html>
  );
}
