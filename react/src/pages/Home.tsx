import reactLogo from "../assets/react.svg";
import viteLogo from "/vite.svg";
// import viteLogo from "";
import { useState } from "react";

export function Home() {
  const [count, setCount] = useState(0);

  return (
    <>
      <SearchInput type="text" />
    </>
  );
}

export function SearchInput({type}) {
  return (<>
  </>)
}
