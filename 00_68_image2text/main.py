import os
from PIL import Image
from transformers import DonutProcessor, VisionEncoderDecoderModel

def extract_text_with_donut(image_path):
    try:
        print(f"Extracting text from image: {image_path}")

        # Donutモデルとプロセッサをロード
        print("Loading Donut model and processor...")
        model = VisionEncoderDecoderModel.from_pretrained("naver-clova-ix/donut-base")
        processor = DonutProcessor.from_pretrained("naver-clova-ix/donut-base")

        # 画像をロード
        print("Loading image...")
        image = Image.open(image_path).convert("RGB")

        # 画像をプロセス
        print("Processing image...")
        pixel_values = processor(image, return_tensors="pt").pixel_values

        # モデルを使用して予測
        print("Generating text from image...")
        decoder_input_ids = processor.tokenizer("<s>", add_special_tokens=False, return_tensors="pt").input_ids
        outputs = model.generate(pixel_values, decoder_input_ids=decoder_input_ids, max_length=512)
        text = processor.batch_decode(outputs, skip_special_tokens=True)[0]

        print(f"Extracted text: {text}")
        return text

    except Exception as e:
        print(f"Error during image processing: {e}")
        return "Error in processing"

def save_text_to_file(text, output_path):
    try:
        with open(output_path, 'w', encoding='utf-8') as file:
            file.write(text)
        print(f"Text saved to file: {output_path}")

    except Exception as e:
        print(f"Error saving text to file: {e}")

def main():
    pre_folder = "image/pre"
    text_folder = "text"

    os.makedirs(text_folder, exist_ok=True)
    print(f"Processing images in folder: {pre_folder}")

    for file_name in os.listdir(pre_folder):
        if file_name.endswith(('.png', '.jpg', '.jpeg','JPG')):
            input_path = os.path.join(pre_folder, file_name)
            file_name_without_ext = os.path.splitext(file_name)[0]
            text_file_path = os.path.join(text_folder, file_name_without_ext + ".txt")

            # Donutで画像からテキスト抽出
            text = extract_text_with_donut(input_path)

            # テキストをファイルに保存
            save_text_to_file(text, text_file_path)

            print(f"Processed and saved text for: {file_name}")

if __name__ == "__main__":
    main()
