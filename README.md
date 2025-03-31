### This is a Golang Ebitenengine Program;

The goals of which are to produce information on the nature of working with ebitengine for tile or grid based games in the hopes that these will be easier to manage than the more free-floating games.

## Visuals:


## Compiling to WebAssembly:
    ```
    go run github.com/hajimehoshi/wasmserve@latest
    ```

    I'm writing this mostly for my future self
    in Windows Powershell enter the code/script below:
    ```
        $Env:GOOS = 'js'
        $Env:GOARCH = 'wasm'
        go build -o bin/GridExperiment.wasm github.com/KelleyTyler/GridTileEbitenDemo03_17
        Remove-Item Env:GOOS
        Remove-Item Env:GOARCH
    ```
    then in the bin folder run:
    ```
        $goroot = go env GOROOT
        cp $goroot\lib\wasm\wasm_exec.js .
    ```
    then in an HTML file have:
    ```
        <!DOCTYPE html>
        <iframe src="main.html" width="640" height="480"></iframe>

        <script src="wasm_exec.js"></script>
        <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("GridExperiment.wasm"), go.importObject).then(result => {
            go.run(result.instance);
        });
        </script>
    ```

    then in VS code run the HTML in 'Live Server' mode



## UPDATE 2025.03.22; a note to self for future work;

    I'm thinking about adding a 2nd layer of "Integer Matrix" to this whole thing;

    reasoning; it would make certain things easier to do; and avoid potentially 'corrupting' data while doing something like Line of Sight calculations on a grid;

    ----
    First the goal should be to get "modular" things out of this build;
    -----

        Rationally this would be some better UI setups: Button handling/etc.. 
        Some kind of "text field entry".. alternatively we could just go make another repo for testing that out;    

