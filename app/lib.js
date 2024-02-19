window.loadGetYaml = () => {
  const go = new Go();
  return fetch('dist/lib.wasm')
    .then((response) => response.arrayBuffer())
    .then((buffer) => WebAssembly.instantiate(buffer, go.importObject))
    .then((result) => {
      go.run(result.instance);

      const getYaml = (templateYaml, valuesYaml) => {
        let returnValueJson;
        try {
          returnValueJson = GetYaml(templateYaml, valuesYaml);
          const returnValue = JSON.parse(returnValueJson);
          return returnValue;
        } catch (error) {
          throw new Error(
            `unable to parse returnValueJson: ${returnValueJson}`
          );
        }
      };

      return getYaml;
    });
};
