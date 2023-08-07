const Sharp = require("sharp");

const main = async () => {
  console.time("Time taken:");

  await processImage({
    watermarkFile: "../watermark.png",
    inputFile: "../input.jpg",
    outputFile: "../node-output.png",
  });

  console.timeEnd("Time taken:");
};

const processImage = async ({ watermarkFile, inputFile, outputFile }) => {
  try {
    await Sharp(inputFile)
      .composite([
        {
          input: watermarkFile,
          top: 64,
          left: 64,
        },
      ])
      .png()
      .toFile(outputFile);
    return { message: "success", fail: false };
  } catch (error) {
    console.error("Failed:", error);
    return { message: error?.message, fail: true };
  }
};

main();
