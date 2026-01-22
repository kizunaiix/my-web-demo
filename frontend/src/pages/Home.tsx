export function Home() {
  return (
    <>
      <main>
        <SearchInput
          inputClassName=""
          buttonClassName="bg-blue-300"
          buttonOnClick={() => {
            window.alert("该跳转搜索结果");
          }}
        />
      </main>
    </>
  );
}

// SearchInput组件

import { cn } from "@/lib/utils";

interface SearchInputProps {
  inputClassName?: string;
  buttonClassName?: string;
  buttonOnClick: () => void;
}

function SearchInput({
  inputClassName,
  buttonClassName,
  buttonOnClick,
}: SearchInputProps) {
  return (
    <>
      <div className="inline-flex relative">
        <input
          type="text "
          className={cn("p-1 border rounded-xl", inputClassName)}
        />
        <button
          onClick={buttonOnClick}
          className={cn(
            "right-0 absolute p-1 border rounded-xl",
            buttonClassName
          )}
        >
          搜索？
        </button>
      </div>
    </>
  );
}

export { SearchInput };
