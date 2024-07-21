'use client';
import { log } from 'console';
import { promises } from 'dns';
import Image from 'next/image'

export default function page() {
  return (
    <div className="flex flex-col min-h-dvh">

      <div className="flex justify-between items-center bg-lime-50/50 mx-auto p-4 w-full">
        <h1 className="font-bold text-xl">logo</h1>
        <nav>
          <ol className='flex space-x-6'>
            <li><a href="/" className="hover:underline">首页</a></li>
            <li><a href="#about" className="hover:underline">历史</a></li>
            <li><a href="/tic-tac-toe" className="hover:underline">收藏</a></li>
            <li><a href="/img/miku_long.jpg" className="hover:underline">自杀</a></li>
          </ol>
        </nav>
      </div>

      <div className="flex flex-grow justify-center">

        <div id="aa" className="md:block relative hidden backdrop-blur-sm w-[360px] fill-transparent">
          <Image
            // width={360}
            // height={600}
            className='hover:blur-none rounded-lg'
            src="/img/喜多ちゃん❤️.jpg"
            layout="fill"
            alt="Picture of 喜多ちゃん"
          />
        </div>

        <div className="flex justify-center bg-slate-100/50">
          <div className="bg-lime-50 pl-10 pr-10 mb-4 w-[40dvw] h-dvh">
            <form action="/" method="post" className="flex flex-col p-1 bg-red-100">
              <div>
                {Input("类型", "text")}
              </div>
              {Input("姓名", "email")}
              <div className='flex justify-center bg-blue-400'>
                {Button("注册", "submit", "indigo", () => { })}
                {Button("登录", "button", "indigo", () => {
                  fetch("http://localhost:8000/api/login", {
                    method: "post",
                    headers: {
                      'Content-Type': 'application/json'
                    },
                    body:JSON.stringify({
                      content: "hello api",
                      content2:"hello!"
                    })
                  })
                  })
                }
              </div>

            </form>
          </div>
        </div>

      </div>
    </div >
  );
}
// const formOnSubmit = ()=>{
//   fetch("localhost:3000")
// }

// function formOnSubmit() {
//   fetch("localhost:3000")
// }

function Button(content: string, type: "submit" | "reset" | "button" | undefined, color: string, fonclick: () => void) {
  return (
    <button type={type} onClick={fonclick} className={`bg-${color}-600 hover:bg-${color}-700 bg shadow-sm m-4 px-4 py-2 border border-transparent rounded-md w-16 font-medium text-sm text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500`}>
      {content}
    </button >
  )
}

function Input(content: string, type: string) {
  return (<div className=''>
    <p>
      {content}
    </p>
    {/* <div className='flex justify-center'> */}
    <input type={type} name="email" size={40} maxLength={50} className="block border-gray-300 focus:border-indigo-500 shadow-sm mt-1 p-2 border rounded-md w-2/3 sm:text-sm focus:outline-none focus:ring-indigo-500" />
    {/* </div> */}
  </div>
  )
}