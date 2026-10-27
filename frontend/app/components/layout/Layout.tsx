import type { ReactNode } from "react"

export function Layout({ children }: { children: ReactNode }) {
  return (
    <>
      <div className="flex flex-col min-h-svh">
        <header className="bg-background h-12">
          <div>这里是页头</div>
        </header>
        <main className="flex-1 bg-background border-border border-y">
          {children}
        </main>
        <footer className="bg-background h-8">
          <div>页脚</div>
        </footer>
      </div>
    </>
  )
}
