export function Home() {
  return (
    <>
      <main>
        <SearchInput
          inputClassName="" // TODO classname的命名是不是要优化
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
import clsx from "clsx";  // TODO clsx没法处理覆盖的问题

interface SearchInputProps {
  inputClassName?: string;
  buttonClassName?: string;
  buttonOnClick: () => void;
}

export function SearchInput({
  inputClassName,
  buttonClassName,
  buttonOnClick,
}: SearchInputProps) {
  return (
    <>
      <div className="inline-flex relative">
        <input
          type="text "
          className={clsx("p-1 border rounded-xl", inputClassName)}
        />
        <button
          onClick={buttonOnClick}
          className={clsx(
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
