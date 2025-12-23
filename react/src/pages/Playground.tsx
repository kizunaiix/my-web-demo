import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

function Playground() {
  return (
    <>
      <main className="flex flex-col content-center bg-background h-screen">
        <div className="flex flex-1 justify-center items-center gap-2 bg-amber-300">

        </div>
        <div className="flex flex-6 justify-center items-center gap-2 bg-amber-400">
          <Input type="text" placeholder="Search" className="w-1/3"/>
          <Button type="submit" variant="default">
            Go
          </Button>
        </div>
      </main>
    </>
  );
}

export { Playground };
