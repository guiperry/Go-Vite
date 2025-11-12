import { Card } from "@/components/ui/card";
import { Copy, Check } from "lucide-react";
import { useState } from "react";
import { Button } from "@/components/ui/button";

interface CodeBlockProps {
  code: string;
  language?: string;
}

export const CodeBlock = ({ code, language = "bash" }: CodeBlockProps) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <Card className="relative bg-card/80 backdrop-blur-sm border-border/50 overflow-hidden group">
      <div className="absolute top-3 right-3">
        <Button
          size="sm"
          variant="ghost"
          onClick={handleCopy}
          className="opacity-0 group-hover:opacity-100 transition-opacity"
        >
          {copied ? <Check className="w-4 h-4" /> : <Copy className="w-4 h-4" />}
        </Button>
      </div>
      <div className="p-4">
        <div className="text-xs text-muted-foreground mb-2 font-mono">{language}</div>
        <pre className="text-sm font-mono text-foreground overflow-x-auto">
          <code>{code}</code>
        </pre>
      </div>
    </Card>
  );
};
