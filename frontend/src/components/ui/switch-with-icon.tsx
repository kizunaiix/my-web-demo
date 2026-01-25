import * as React from "react"
import { Switch as SwitchPrimitive } from "radix-ui"
import { Moon, Sun } from "lucide-react"
import { cn } from "@/lib/utils"

interface SwitchWithIconProps extends Omit<React.ComponentProps<typeof SwitchPrimitive.Root>, 'children'> {
  size?: "sm" | "default";
  icons?: {
    checked: React.ReactNode;
    unchecked: React.ReactNode;
  };
  defaultIcons?: "dark-mode" | null;
}

function SwitchWithIcon({ 
  checked, 
  onCheckedChange, 
  size = "default",
  icons,
  defaultIcons = null,
  className,
  ...props 
}: SwitchWithIconProps) {
  // 图标选择逻辑
  const selectedIcons = React.useMemo(() => {
    if (icons) return icons;
    if (defaultIcons === "dark-mode") {
      return {
        checked: <Moon />,
        unchecked: <Sun />
      };
    }
    return null;
  }, [icons, defaultIcons]);

  // 图标尺寸计算 (95%比例)
  const iconSize = size === "default" ? 15 : 11;
  
  // 如果没有图标，回退到普通Switch
  if (!selectedIcons) {
    return (
      <SwitchPrimitive.Root
        checked={checked}
        onCheckedChange={onCheckedChange}
        className={cn(
          // 复用原Switch样式
          "data-checked:bg-primary data-unchecked:bg-input focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:aria-invalid:border-destructive/50 dark:data-unchecked:bg-input/80 shrink-0 rounded-full border border-transparent shadow-xs focus-visible:ring-[3px] aria-invalid:ring-[3px] data-[size=default]:h-[18.4px] data-[size=default]:w-[32px] data-[size=sm]:h-[14px] data-[size=sm]:w-[24px] peer group/switch relative inline-flex items-center transition-all outline-none after:absolute after:-inset-x-3 after:-inset-y-2 data-disabled:cursor-not-allowed data-disabled:opacity-50",
          className
        )}
        {...props}
      >
        <SwitchPrimitive.Thumb
          data-slot="switch-thumb"
          className="bg-background dark:data-unchecked:bg-foreground dark:data-checked:bg-primary-foreground rounded-full group-data-[size=default]/switch:size-4 group-data-[size=sm]/switch:size-3 group-data-[size=default]/switch:data-checked:translate-x-[calc(100%-2px)] group-data-[size=sm]/switch:data-checked:translate-x-[calc(100%-2px)] group-data-[size=default]/switch:data-unchecked:translate-x-0 group-data-[size=sm]/switch:data-unchecked:translate-x-0 pointer-events-none block ring-0 transition-transform"
        />
      </SwitchPrimitive.Root>
    );
  }

  return (
    <SwitchPrimitive.Root
      data-slot="switch-with-icon"
      data-size={size}
      checked={checked}
      onCheckedChange={onCheckedChange}
      className={cn(
        // 继承原Switch样式
        "data-checked:bg-primary data-unchecked:bg-input focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:aria-invalid:border-destructive/50 dark:data-unchecked:bg-input/80 shrink-0 rounded-full border border-transparent shadow-xs focus-visible:ring-[3px] aria-invalid:ring-[3px] data-[size=default]:h-[18.4px] data-[size=default]:w-[32px] data-[size=sm]:h-[14px] data-[size=sm]:w-[24px] peer group/switch relative inline-flex items-center transition-all outline-none after:absolute after:-inset-x-3 after:-inset-y-2 data-disabled:cursor-not-allowed data-disabled:opacity-50",
        className
      )}
      {...props}
    >
      <SwitchPrimitive.Thumb
        data-slot="switch-with-icon-thumb"
        className={cn(
          "bg-background dark:data-unchecked:bg-foreground dark:data-checked:bg-primary-foreground rounded-full group-data-[size=default]/switch:size-4 group-data-[size=sm]/switch:size-3 group-data-[size=default]/switch:data-checked:translate-x-[calc(100%-2px)] group-data-[size=sm]/switch:data-checked:translate-x-[calc(100%-2px)] group-data-[size=default]/switch:data-unchecked:translate-x-0 group-data-[size=sm]/switch:data-unchecked:translate-x-0 pointer-events-none block ring-0 transition-transform flex items-center justify-center overflow-hidden",
        )}
      >
        {/* 图标容器 */}
        <div className="relative w-full h-full flex items-center justify-center">
          {/* 未选中状态图标 */}
          <div 
            className={cn(
              "absolute transition-all duration-200 flex items-center justify-center",
              size === "default" ? "w-[15px] h-[15px]" : "w-[11px] h-[11px]",
              checked 
                ? "opacity-0 scale-0 rotate-180" 
                : "opacity-100 scale-100 rotate-0 text-foreground"
            )}
          >
            {selectedIcons.unchecked}
          </div>
          
          {/* 选中状态图标 */}
          <div 
            className={cn(
              "absolute transition-all duration-200 flex items-center justify-center",
              size === "default" ? "w-[15px] h-[15px]" : "w-[11px] h-[11px]",
              checked 
                ? "opacity-100 scale-100 rotate-0 text-foreground" 
                : "opacity-0 scale-0 -rotate-180 text-foreground"
            )}
          >
            {selectedIcons.checked}
          </div>
        </div>
      </SwitchPrimitive.Thumb>
    </SwitchPrimitive.Root>
  )
}

export { SwitchWithIcon }