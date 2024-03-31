
import * as monaco from 'monaco-editor';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';
import { requiem } from "./defs.json";
/** define MonacoEnvironment.getWorker  */
(self as any).MonacoEnvironment = {
    getWorker(_: any, label: string) {
        if (label === 'typescript' || label === 'javascript') {
            return new tsWorker();
        }
        return new editorWorker();
    }
};
monaco.languages.typescript.typescriptDefaults.setEagerModelSync(true);
console.log(requiem);
monaco.languages.typescript.typescriptDefaults.addExtraLib(requiem)

