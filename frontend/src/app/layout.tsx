import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'mDawka',
    description: 'Jaka piguła wariacie?',
}

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
            <body>{children}</body>
        </html>
    )
}
