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
        <Header isDark={isDark} setIsDark={setIsDark} />
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

// ---
import { SwitchWithIcon } from "@/components/ui/switch-with-icon";
interface HeaderProps {
  isDark: boolean;
  setIsDark: React.Dispatch<React.SetStateAction<boolean>>;
}
function Header({ isDark, setIsDark }: HeaderProps) {
  return (
    <>
      <div className="flex items-center gap-2 p-4">
        <SwitchWithIcon
          checked={isDark}
          onCheckedChange={setIsDark}
          defaultIcons="dark-mode"
          className="cursor-pointer"
        />
        <span className="text-muted-foreground text-sm">
          {isDark ? 'Dark' : 'Light'}
        </span>
      </div>
    </>
  );
}
