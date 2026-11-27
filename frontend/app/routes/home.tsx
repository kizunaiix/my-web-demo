import { useState } from "react"
import { Button } from "~/components/ui/button"

export default function Home() {
  const [count, setCount] = useState(1)
  return (
    <div className="flex p-6 min-h-svh">
      <div className="flex flex-col gap-4 min-w-0 max-w-md text-sm leading-loose">
        <div>
          <h1 className="font-medium">Project ready!</h1>
          <p>You may now add components and start building.</p>
          <p>We&apos;ve already added the button component for you.</p>
          <Button className="mt-2" onClick={() => setCount(c => c + 1)}>
            Button {count}
          </Button>
        </div>
      </div>
    </div>
  )
}
