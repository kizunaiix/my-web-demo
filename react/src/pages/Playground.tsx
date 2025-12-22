import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

function Playground() {
  return (
    <>
      <main className="flex bg-background h-screen">
        <div className="flex justify-center items-center gap-2 w-full">
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
