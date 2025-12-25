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
        <div className="bg-red-300 w-12 h-64">
          ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss
        </div>
        <div className="inline-block bg-blue-600 h-64">
          <div className="inline bg-red-500 border w-6 h-32">
            12dfsfsdfdsfdsfsd
          </div>
          <div className="inline bg-red-500 border w-6 h-32">2</div>
          <div className="bg-red-500 border w-6 h-32"></div>
        </div>
      </main>
    </>
  );
}

// SearchInput组件

import { cn } from "@/lib/utils"

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
          搜索
        </button>
      </div>
    </>
  );
}

export { SearchInput };
