import { useState, useEffect } from "react";

interface LayoutProps {
  withSidebar: boolean;
  children: React.ReactNode;
}

function Layout({ withSidebar, children }: LayoutProps) {
  const [isDark, setIsDark] = useState(false);

  useEffect(() => {
    document.documentElement.classList.toggle("dark", isDark);
  }, [isDark]);

  return (
    <>
      <div className="flex flex-col min-h-dvh">
        <header>header!</header>
        {/*  TODO 把isDark传给header用于暗黑模式切换按钮 */}
        <div className="flex flex-1">
          {withSidebar && (
            <aside className="hidden md:block bg-sidebar border w-32">
              sidebar
            </aside>
          )}
          <main className="flex-1">{children}</main>
        </div>
        <footer>thefooter</footer>
      </div>
    </>
  );
}

export { Layout };