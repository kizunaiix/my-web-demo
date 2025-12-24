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
        <div className="bg-red-300 w-12 h-64">
          ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss
        </div>
        <div className="inline-block bg-blue-600 h-64">
          <div className="inline bg-red-500 border w-6 h-32">12dfsfsdfdsfdsfsd</div>
          <div className="inline bg-red-500 border w-6 h-32">2</div>
          <div className="bg-red-500 border w-6 h-32"></div>
        </div>
      </main>
    </>
  );
}

// SearchInput组件
import clsx from "clsx"; // TODO clsx没法处理覆盖的问题

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
