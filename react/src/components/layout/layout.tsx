// import { useState } from "react";

interface layoutProps {
  withSidebar: boolean;
  children: React.ReactNode;
}

function Layout({ withSidebar, children }: layoutProps) {
  return (
    <>
      <div className="flex flex-col min-h-dvh">
        <header>header!</header>
        <div className="flex flex-1">
          {withSidebar && <aside className="hidden md:block w-32">sidebar</aside>}
          <main className="flex-1">{children}</main>
        </div>
        <footer>thefooter</footer>
      </div>
    </>
  );
}

export { Layout };
