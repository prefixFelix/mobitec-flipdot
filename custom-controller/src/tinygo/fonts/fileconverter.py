def extract_font_data(header_file_content, output_file_path):
    """
    Extract font byte data from a C/C++ header file and save it as a binary file.

    Args:
        header_file_content (str): Content of the header file
        output_file_path (str): Path where the binary file will be saved
    """
    # Split content by lines
    lines = header_file_content.strip().split('\n')

    # Find the line that contains the array declaration
    array_start_index = -1
    for i, line in enumerate(lines):
        if "PROGMEM" in line and "[" in line and "]" in line:
            array_start_index = i + 1
            break

    if array_start_index == -1:
        raise ValueError("Could not find the array declaration in the header file")

    # Extract bytes from the data section
    bytes_data = []

    # Process each line after the array declaration
    for line in lines[array_start_index:]:
        # Skip empty lines and comments
        if not line.strip() or line.strip().startswith('//'):
            print(line)
            continue

        # Stop processing when we reach the end of the array
        if line.strip() == "};":
            print(line)
            break

        # Extract hex values from the line
        parts = line.split(',')
        for part in parts:
            # Clean up the part to extract just the hex value
            cleaned_part = part.strip()
            if not cleaned_part or cleaned_part == "":
                print(cleaned_part)
                continue

            try:
                # Convert hex (0xNN) or decimal values to integer
                if cleaned_part.startswith('0x'):
                    value = int(cleaned_part, 16)
                else:
                    value = int(cleaned_part)

                # Ensure the value is a valid byte (0-255)
                if 0 <= value <= 255:
                    bytes_data.append(value)
            except ValueError:
                # Skip non-numeric entries
                continue

    # Convert the list of integers to bytes
    binary_data = bytes(bytes_data)

    # Write bytes to binary file
    with open(output_file_path, 'wb') as f:
        f.write(binary_data)

    print(f"Font data extracted and saved to {output_file_path}")
    print(f"Extracted {len(binary_data)} bytes")

# Read the header file
def main():
    input_file_path = "wwFont_Xx7_v02.h"
    output_file_path = "wwXx7.bin"

    try:
        with open(input_file_path, 'r') as f:
            header_content = f.read()

        extract_font_data(header_content, output_file_path)
    except FileNotFoundError:
        print(f"Error: Could not find the file {input_file_path}")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()
